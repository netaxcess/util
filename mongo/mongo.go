//Mongodb
package mongo

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	"time"
)

type Mongodb struct {
	*mgo.Session
	Dsn			string
	Database    string
}

/*
例子：mongodb, err := NewMongodb("mongodb://ichunt:huntmon6699@192.168.1.237/ichunt", "ichunt")
*/

func NewMongodb(dsn string, database string) (*Mongodb, error) {
	m := &Mongodb{
		Dsn		  : dsn,
		Database  : db,
	}
	return m, m.connect()
}

func (m *Mongodb) connect() error {
	session, err := mgo.DialWithTimeout(m.Dsn, 10 * time.Second)
	if err != nil {
		return err
	}

	m.Session = session
	return nil
}

func (m *Mongodb) WithFunc(collection string, fs func(*mgo.Collection) error) error {
	s := m.Session.New()
	err := fs(s.DB(m.Database).C(collection))
	s.Close()
	return err
}

func (self *Mongodb) Insert(collection string, data ...interface{}) error {
	return self.WithFunc(collection, func(c *mgo.Collection) error {
		return c.Insert(data...)
	})
}

/*
例子：
	latestLog := &sku{
		stock: 1,
		moq  : 1,
	}
	Update("sku", bson.M{"sku_name": "lm358", "supplier_id": 2}, latestLog)

	goods := new(Mgoods)
	 mongodb.Update("chip1stop", bson.M{"goods_id": 12000148992}, goods)
*/
func (self *Mongodb) Update(collection string, selector interface{}, change interface{}) error {
	return self.WithFunc(collection, func(c *mgo.Collection) error {
		_, err := c.Upsert(selector, change)
		return err
	})
}

/*
例子：FindOne("sku", bson.M{"sku_id": 100}, &sku_struct)
*/
func (self *Mongodb) FindOne(collection string, query interface{}, result interface{}) error {
	return self.WithFunc(collection, func(c *mgo.Collection) error {
		return c.Find(query).One(result)
	})
}

/*
例子：Delete("sku",bson.M{"goods_id":sku_id})
*/
func (self *Mongodb) Delete(collection string, args bson.M) error {
	return self.WithFunc(collection, func(c *mgo.Collection) error {
		_, err := c.RemoveAll(args)
		return err
	})
}


