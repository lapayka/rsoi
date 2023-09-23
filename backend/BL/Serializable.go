package BL

type Serializable interface {
	ToJSON() string
	FromJSON(string) error
}
