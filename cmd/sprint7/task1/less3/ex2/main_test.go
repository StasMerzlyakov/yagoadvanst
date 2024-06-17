package ex2_test

// JSONData — интерфейс для декодирования JSON.
type JSONData interface {
	DecodeJSON() interface{}
}

// YAMLData — интерфейс для декодирования YAML.
type YAMLData interface {
	DecodeYAML() interface{}
}

type Client struct {
	Data interface{}
}

func (client *Client) Decode(input JSONData) {
	client.Data = input.DecodeJSON()
}

type Adapter struct {
	YAMLData YAMLData
}

func (adpt *Adapter) DecodeJSON() interface{} {
	return adpt.YAMLData.DecodeYAML()
}

// добавьте тип Adapter и необходимый метод
// ...

func Load(client Client, input YAMLData) {
	adapter := &Adapter{
		YAMLData: input,
	}
	client.Decode(adapter)
}
