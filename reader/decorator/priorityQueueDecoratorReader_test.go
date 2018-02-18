package decorator

import (
	"testing"
	"github.com/ufoscout/go-up/reader"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldBeAssignableToReader(t *testing.T) {

	var prop reader.Reader = &PriorityQueueDecoratorReader{map[int][]reader.Reader{}}

	result, err := prop.Read();
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func Test_ShouldReturnEmptyMapIfEmpty(t *testing.T) {
	queue := NewPriorityQueueDecoratorReader();
	prop, _ := queue.Read();
	assert.NotNil(t, prop);
	assert.True(t, len(prop) == 0);
}

func Test_ShouldMergeEntriesFromMapWithDifferentPriority(t *testing.T) {
	queue := NewPriorityQueueDecoratorReader();
	queue.Add(reader.NewProgrammaticReader().Add("k1", "v1").Add("k2", "v2"), 11);
	queue.Add(reader.NewProgrammaticReader().Add("k3", "v3").Add("k4", "v4"), 1);
	prop, _ := queue.Read();
	assert.NotNil(t, prop);
	assert.Equal(t, 4, len(prop));
	assert.Equal(t, "v1", prop["k1"].Value);
	assert.Equal(t, "v2", prop["k2"].Value);
	assert.Equal(t, "v3", prop["k3"].Value);
	assert.Equal(t, "v4", prop["k4"].Value);
}

func Test_ShouldMergeEntriesFromMapWithSamePriority(t *testing.T) {
	queue := NewPriorityQueueDecoratorReader();
	queue.Add(reader.NewProgrammaticReader().Add("k1", "v1").Add("k2", "v2"), 11);
	queue.Add(reader.NewProgrammaticReader().Add("k3", "v3").Add("k4", "v4"), 11);
	prop, _ := queue.Read();
	assert.NotNil(t, prop);
	assert.Equal(t, 4, len(prop));
	assert.Equal(t, "v1", prop["k1"].Value);
	assert.Equal(t, "v2", prop["k2"].Value);
	assert.Equal(t, "v3", prop["k3"].Value);
	assert.Equal(t, "v4", prop["k4"].Value);
}

func Test_ShouldTakeIntoAccountPriorityInCaseOfCollisions(t *testing.T) {
	queue := NewPriorityQueueDecoratorReader();
	queue.Add(reader.NewProgrammaticReader().Add("k1", "v1").Add("k2", "v2-first"), 2);
	queue.Add(reader.NewProgrammaticReader().Add("k3", "v3").Add("k2", "v2-second"), 1);
	prop, _ := queue.Read();
	assert.NotNil(t, prop);
	assert.Equal(t, 3, len(prop));
	assert.Equal(t, "v1", prop["k1"].Value);
	assert.Equal(t, "v2-second", prop["k2"].Value);
	assert.Equal(t, "v3", prop["k3"].Value);
}

func Test_ShouldTakeIntoAccountInsertionOrderForSamePriorityInCaseOfCollisions(t *testing.T) {
	queue := NewPriorityQueueDecoratorReader();
	queue.Add(reader.NewProgrammaticReader().Add("k3", "v3").Add("k2", "v2-second"), 1);
	queue.Add(reader.NewProgrammaticReader().Add("k1", "v1").Add("k2", "v2-first"), 1);
	prop, _ := queue.Read();
	assert.NotNil(t, prop);
	assert.Equal(t, 3, len(prop));
	assert.Equal(t, "v1", prop["k1"].Value);
	assert.Equal(t, "v2-first", prop["k2"].Value);
	assert.Equal(t, "v3", prop["k3"].Value);
}

func Test_ShouldTakeIntoAccountInsertionOrderAndPriority(t *testing.T) {
	queue := NewPriorityQueueDecoratorReader();
	queue.Add(reader.NewProgrammaticReader().Add("k2", "v2-third").Add("k3", "v3-third").Add("k5", "v5-third"), 5);
	queue.Add(reader.NewProgrammaticReader().Add("k1", "v1-first").Add("k2", "v2-first").Add("k4", "v4-first"), 10);
	queue.Add(reader.NewProgrammaticReader().Add("k1", "v1-second").Add("k2", "v2-second").Add("k3", "v3-second"), 10);
	prop, _ := queue.Read();
	assert.NotNil(t, prop);
	assert.Equal(t, 5, len(prop));
	assert.Equal(t, "v1-second", prop["k1"].Value);
	assert.Equal(t, "v2-third", prop["k2"].Value);
	assert.Equal(t, "v3-third", prop["k3"].Value);
	assert.Equal(t, "v4-first", prop["k4"].Value);
	assert.Equal(t, "v5-third", prop["k5"].Value);
}

func Test_ShouldReturnOsEnv(t *testing.T) {
	queue := NewPriorityQueueDecoratorReader()
	queue.Add(&reader.EnvReader{""}, 0)
	prop, _ := queue.Read();
	assert.NotNil(t, prop);
	assert.True(t, len(prop) > 0);
}
