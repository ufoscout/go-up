package go_up

import (
	"github.com/ufoscout/go-up/reader"
	"github.com/ufoscout/go-up/reader/decorator"
)

const HIGHEST_PRIORITY int = 0
const DEFAULT_PRIORITY int = 100

const DEFAULT_START_DELIMITER string = "${"
const DEFAULT_END_DELIMITER string = "}"

const DEFAULT_LIST_SEPARATOR string = ","

type GoUpBuilder interface {
	Add(key string, value string) GoUpBuilder
	AddReader(newReader reader.Reader) GoUpBuilder
	AddReaderWithPriority(newReader reader.Reader, priority int) GoUpBuilder
	AddFile(filename string, ignoreNotFound bool) GoUpBuilder
	AddFileWithPriority(filename string, ignoreNotFound bool, priority int) GoUpBuilder
	Delimiters(startDelimiter string, endDelimiter string) GoUpBuilder
	IgnoreUnresolvablePlaceholders(ignoreUnresolvablePlaceholders bool) GoUpBuilder
	Build() (GoUp, error)
}

func NewGoUp() GoUpBuilder {
	return &goUpBuilderImpl{decorator.NewPriorityQueueDecoratorReader(),
		DEFAULT_START_DELIMITER,
		DEFAULT_END_DELIMITER,
		false}
}

type goUpBuilderImpl struct {
	reader                         *decorator.PriorityQueueDecoratorReader
	startDelimiter                 string
	endDelimiter                   string
	ignoreUnresolvablePlaceholders bool
}

func (up *goUpBuilderImpl) Add(key string, value string) GoUpBuilder {
	return up.AddReaderWithPriority(reader.NewProgrammaticReader().Add(key, value), DEFAULT_PRIORITY)
}

/**
 * Add a new property Reader with the default priority.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 *
 */
func (up *goUpBuilderImpl) AddReader(newReader reader.Reader) GoUpBuilder {
	return up.AddReaderWithPriority(newReader, DEFAULT_PRIORITY)
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
	return up.AddFileWithPriority(filename, ignoreNotFound, DEFAULT_PRIORITY)
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
 * Whether to ignore not resolvable placeholders.
 * Default is false.
 */
func (up *goUpBuilderImpl) IgnoreUnresolvablePlaceholders(ignoreUnresolvablePlaceholders bool) GoUpBuilder {
	up.ignoreUnresolvablePlaceholders = ignoreUnresolvablePlaceholders
	return up
}

func (up *goUpBuilderImpl) Build() (GoUp, error) {

	replacer := decorator.PlaceholderReplacerDecoratorReader{up.reader, up.startDelimiter, up.endDelimiter, up.ignoreUnresolvablePlaceholders}

	properties, err := replacer.Read()

	if err != nil {
		return nil, err
	}

	return &goUpImpl{properties}, nil
}
