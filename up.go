package go_up

import (
	"github.com/ufoscout/go-up/reader/decorator"
	"github.com/ufoscout/go-up/reader"
)

const HIGHEST_PRIORITY int = 0
const DEFAULT_PRIORITY int = 100

const DEFAULT_START_DELIMITER string = "${"
const DEFAULT_END_DELIMITER string = "}"

const DEFAULT_LIST_SEPARATOR string = ",";

type GoUp interface {
}

func NewGoUp() GoUp {
	return &goUpImpl{decorator.NewPriorityQueueDecoratorReader(),
		DEFAULT_START_DELIMITER,
		DEFAULT_END_DELIMITER,
		false}
}

type goUpImpl struct {
	reader                         *decorator.PriorityQueueDecoratorReader
	startDelimiter                 string
	endDelimiter                   string
	ignoreUnresolvablePlaceholders bool
}

/**
 * Add a new property Reader with the default priority.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 *
 */
func (up *goUpImpl) AddReader(newReader reader.Reader) GoUp {
	return up.AddReaderWithPriority(newReader, DEFAULT_PRIORITY)
}

/**
 * Add a new property Reader.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 *
 */
func (up *goUpImpl) AddReaderWithPriority(newReader reader.Reader, priority int) GoUp {
	up.reader.Add(up.reader, priority)
	return up
}

/**
 * Add a new properties file with the default priority.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 */
func (up *goUpImpl) AddFile(filename string, ignoreNotFound bool) GoUp {
	return up.AddFileWithPriority(filename, ignoreNotFound, DEFAULT_PRIORITY)
}

/**
 * Add a new properties file with the specified priority.
 * If two or more Readers have the same priority, the last added has the highest priority among them.
 */
func (up *goUpImpl) AddFileWithPriority(filename string, ignoreNotFound bool, priority int) GoUp {
	return up.AddReaderWithPriority(&reader.FileReader{filename, ignoreNotFound}, priority)
}

/**
 *
 * Set the start and end placeholder delimiters.
 * Default are {@value Default#START_DELIMITER} and {@value Default#END_DELIMITER}
 */
func (up *goUpImpl) Delimiters(startDelimiter string, endDelimiter string) GoUp {
	up.startDelimiter = startDelimiter
	up.endDelimiter = endDelimiter
	return up
}

/**
 * Whether to ignore not resolvable placeholders.
 * Default is false.
 */
func (up *goUpImpl) IgnoreUnresolvablePlaceholders(ignoreUnresolvablePlaceholders bool) GoUp {
	up.ignoreUnresolvablePlaceholders = ignoreUnresolvablePlaceholders
	return up
}
