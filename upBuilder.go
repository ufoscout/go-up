package goup

import (
	"github.com/ufoscout/go-up/reader"
	"github.com/ufoscout/go-up/reader/decorator"
)

// HighestPriority the highest priority
const HighestPriority int = 0

// DefaultPriority the default priority
const DefaultPriority int = 100

// DefaultStartDelimiter the default placeholder start delimiter
const DefaultStartDelimiter string = "${"

// DefaultEndDelimiter the default placeholder end delimiter
const DefaultEndDelimiter string = "}"

// DefaultValueSeparator the default separator for a placeholder default value
const DefaultValueSeparator string = ":"

// DefaultListSeparator default list separator
const DefaultListSeparator string = ","

// GoUpBuilder the GoUp builder
type GoUpBuilder interface {
	Add(key string, value string) GoUpBuilder
	AddReader(newReader reader.Reader) GoUpBuilder
	AddReaderWithPriority(newReader reader.Reader, priority int) GoUpBuilder
	AddFile(filename string, ignoreNotFound bool) GoUpBuilder
	AddFileWithPriority(filename string, ignoreNotFound bool, priority int) GoUpBuilder
	Delimiters(startDelimiter string, endDelimiter string) GoUpBuilder
	DefaultValueSeparator(defaultValueSeparator string) GoUpBuilder
	IgnoreUnresolvablePlaceholders(ignoreUnresolvablePlaceholders bool) GoUpBuilder
	Build() (GoUp, error)
}

// NewGoUp creates a new GoUpBuilder
func NewGoUp() GoUpBuilder {
	return &goUpBuilderImpl{decorator.NewPriorityQueueDecoratorReader(),
		DefaultStartDelimiter,
		DefaultEndDelimiter,
		DefaultValueSeparator,
		false}
}

type goUpBuilderImpl struct {
	reader                         *decorator.PriorityQueueDecoratorReader
	startDelimiter                 string
	endDelimiter                   string
	defaultValueSeparator          string
	ignoreUnresolvablePlaceholders bool
}

func (up *goUpBuilderImpl) Add(key string, value string) GoUpBuilder {
	return up.AddReaderWithPriority(reader.NewProgrammaticReader().Add(key, value), DefaultPriority)
}

/**
 * Add a new property Reader with the default priority.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 *
 */
func (up *goUpBuilderImpl) AddReader(newReader reader.Reader) GoUpBuilder {
	return up.AddReaderWithPriority(newReader, DefaultPriority)
}

/**
 * Add a new property Reader.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 *
 */
func (up *goUpBuilderImpl) AddReaderWithPriority(newReader reader.Reader, priority int) GoUpBuilder {
	up.reader.Add(newReader, priority)
	return up
}

/**
 * Add a new properties file with the default priority.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 */
func (up *goUpBuilderImpl) AddFile(filename string, ignoreNotFound bool) GoUpBuilder {
	return up.AddFileWithPriority(filename, ignoreNotFound, DefaultPriority)
}

/**
 * Add a new properties file with the specified priority.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 */
func (up *goUpBuilderImpl) AddFileWithPriority(filename string, ignoreNotFound bool, priority int) GoUpBuilder {
	return up.AddReaderWithPriority(&reader.FileReader{filename, ignoreNotFound}, priority)
}

/**
 *
 * Set the start and end placeholder delimiters.
 * Default are {@value Default#START_DELIMITER} and {@value Default#END_DELIMITER}
 */
func (up *goUpBuilderImpl) Delimiters(startDelimiter string, endDelimiter string) GoUpBuilder {
	up.startDelimiter = startDelimiter
	up.endDelimiter = endDelimiter
	return up
}

/**
 *
 * Set the separator for the default value.
 * Default is ':'
 */
func (up *goUpBuilderImpl) DefaultValueSeparator(defaultValueSeparator string) GoUpBuilder {
	up.defaultValueSeparator = defaultValueSeparator
	return up
}

/**
 * Whether to ignore not resolvable placeholders.
 * Default is false.
 */
func (up *goUpBuilderImpl) IgnoreUnresolvablePlaceholders(ignoreUnresolvablePlaceholders bool) GoUpBuilder {
	up.ignoreUnresolvablePlaceholders = ignoreUnresolvablePlaceholders
	return up
}

func (up *goUpBuilderImpl) Build() (GoUp, error) {

	replacer := decorator.PlaceholderReplacerDecoratorReader{
		up.reader,
		up.startDelimiter,
		up.endDelimiter,
		up.defaultValueSeparator,
		up.ignoreUnresolvablePlaceholders}

	properties, err := replacer.Read()

	if err != nil {
		return nil, err
	}

	return &goUpImpl{properties}, nil
}
