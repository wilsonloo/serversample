package models

// Client 客户端信息
type Client struct {
	ID     string `bson:"_id"`
	Name   string `bson:"Name"`
	Secret string `bson:"Secret"`
	Domain string `bson:"Domain"`
}

// CName 存储的集合名称
func (c *Client) CName() string {
	return "Client"
}

// Create 创建客户端信息
func (c *Client) Create() error {
	err := mHandler.C(c.CName()).RemoveId(c.ID)
	if err != nil {
		return err
	}
	return mHandler.C(c.CName()).Insert(c)
}

// GetByID 根据ID获取客户端信息
func (c *Client) GetByID(id string) (*Client, error) {
	var result Client
	err := mHandler.C(c.CName()).FindId(id).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
