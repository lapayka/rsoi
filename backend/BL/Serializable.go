package BL

type bytes = []byte

type Serializable interface {
	ToJSON() bytes
	FromJSON(bytes) any
}
