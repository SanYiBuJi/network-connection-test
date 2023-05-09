package Internal

type DataCache struct {
	data map[string]interface{}
}

func NewDataCache() *DataCache {
	return &DataCache{data: make(map[string]interface{})}
}

func (c *DataCache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *DataCache) Get(key string) interface{} {
	if val, ok := c.data[key]; ok {
		return val
	}
	return nil
}
