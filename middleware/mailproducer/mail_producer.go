package mailproducer

import (
	"known-anchors/dal/mq"
	"sync"

	"known-anchors/util/close"

	"github.com/gin-gonic/gin"
)

var once sync.Once

var mailproducer *mq.Producer

func MailProducerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		once.Do(func() {
			mailproducer = mq.NewProducer("mail")
			go mailproducer.Produce()
			cl := c.MustGet("closechan").(*close.Close)
			cl.AddCloseable(mailproducer)
		})

		c.Set("mailproducer", mailproducer.ProChan)
		c.Next()
	}
}
