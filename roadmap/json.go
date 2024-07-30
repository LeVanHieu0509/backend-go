package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Bird struct {
	Species     string
	Description string
}

/*
Khi convert từ restApi thì cần phải dịnh nghĩa lại api đó thành các struct ở trong go
Và lúc đó trong go sẽ get data thông qua struct
*/
func main() {
	DecodeJsonToStruct()
	fmt.Println("------------------------")
	JsonArray()
	fmt.Println("------------------------")
	NestObject()
	fmt.Println("------------PrimitiveTypes------------")
	PrimitiveTypes()
	fmt.Println("------------TimeValues------------")
	TimeValues()
	fmt.Println("------------CustomParsingLogic------------")
	data := []byte("\"20x30\"")
	var d Dimensions
	err := d.CustomParsingLogic(data)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Height: %d, Width: %d\n", d.Height, d.Width)
	}
	fmt.Println()
	fmt.Println("------------JSONStructTags------------")
	JSONStructTags()
	fmt.Println()
	fmt.Println("------------DecodingJSONtoMaps------------")
	DecodingJSONtoMaps()
	fmt.Println()
	fmt.Println("------------ValidateJsonData------------")
	ValidateJsonData()
	fmt.Println()
	fmt.Println("------------MarshalingJSONData------------")
	MarshalingJSONData()
	fmt.Println()
	fmt.Println("------------IgnoringEmptyFields------------")
	IgnoringEmptyFields()
	fmt.Println()
	fmt.Println("------------MarshalingSlices------------")
	MarshalingSlices()
	fmt.Println()
	fmt.Println("------------MarshalingMaps------------")
	MarshalingMaps()
	fmt.Println()
	fmt.Println("------------EncodingNullValues------------")
	EncodingNullValues()
	fmt.Println()

	fmt.Println("------------CustomEncodingLogic------------")
	CustomEncodingLogic()
	fmt.Println()

	fmt.Println("------------PrintingFormattedPrettyPrintedJSON------------")
	PrintingFormattedPrettyPrintedJSON()
	fmt.Println()

	fmt.Println("------------PracticeJsonDynamic------------")
	PracticeJsonDynamic()
	fmt.Println()

}

// Decoding JSON Into Structs use Unmarshal
var birdJson = `{"species": "pigeon","description": "likes to perch on rocks"}`
var bird Bird

func DecodeJsonToStruct() {

	//bind qua struct bird => khi sài thì lấy ra sài như 1 struct
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)

}

func JsonArray() {

	birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds : %+v\n", birds)

	for _, el := range birds {
		fmt.Println("el: ", el.Species+":"+el.Description)

	}
}

type Dimensions struct {
	Height int
	Width  int
}

type Bird1 struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func NestObject() {
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var birds Bird1
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Println(birds)
}

func PrimitiveTypes() {
	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n)
	// 3

	json.Unmarshal([]byte(floatJson), &pi)
	fmt.Println(pi)
	// 3.1412

	json.Unmarshal([]byte(stringJson), &str)
	fmt.Println(str)
	// bird
}

type Bird2 struct {
	Species     string
	Description string
	CreatedAt   time.Time
}

func TimeValues() {
	dateJson := `"2021-10-18T11:08:47.577Z"`
	var date time.Time
	json.Unmarshal([]byte(dateJson), &date)

	fmt.Println(date)

	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "createdAt": "2021-10-18T11:08:47.577Z"}`
	var bird Bird2
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
}

type Dimensions1 struct {
	Height int
	Width  int
}

/*
Thực hiện logic tùy chỉnh cho việc phân tích cú pháp một chuỗi byte thành một cấu trúc Dimensions
*/
func (d *Dimensions) CustomParsingLogic(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("dimensions string too short")
	}
	// remove the quotes
	s := string(data)[1 : len(data)-1]
	// split the string into its two parts
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return fmt.Errorf("dimensions string must contain two parts")
	}
	// convert the two parts into ints
	height, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("dimension height must be an int")
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("dimension width must be an int")
	}
	// assign the two ints to the Dimensions struct
	d.Height = height
	d.Width = width
	return nil
}

// JSON Struct Tags - Custom Field Names
type Bird3 struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func JSONStructTags() {
	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird3
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
}

func DecodingJSONtoMaps() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]any type, and its type is asserted from
	// the `any` type
	birds := result["birds"].(map[string]any)

	for key, value := range birds {
		// Each value is an `any` type, that is type asserted as a string
		fmt.Println(key+":", value.(string))
	}
}

func ValidateJsonData() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`
	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)

	if !json.Valid([]byte(birdJson)) {
		// handle the error here
		fmt.Println("invalid JSON string:", birdJson)
		return
	}
}

type Bird4 struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

// Marshaling là quá trình chuyển đổi dữ liệu có cấu trúc thành chuỗi JSON có thể tuần tự hóa.
// Tương tự như giải mã, chúng ta có thể giải mã dữ liệu thành các struct, bản đồ và lát cắt.
func MarshalingJSONData() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marshal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
}

type Bir5 struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does,omitempty"`
	Species1    string `json:"-"` // ignore field Species1 in struct
}

// Muốn bỏ qua một trường trong đầu ra JSON của mình, nếu giá trị của trường đó trống. Chúng ta có thể sử dụng thuộc tính "omitempty" cho mục đích này.
func IgnoringEmptyFields() {
	var c Bir5

	// ignore field Description
	pigeon := &Bir5{
		Species:  "Pigeon",
		Species1: "Pigeon",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(string(data))
	// json.Unmarshal([]byte(birdJson), &result)

	json.Unmarshal(data, &c)
	fmt.Println("Struct:", c)

}

// Chúng ta chỉ cần truyền lát cắt hoặc mảng cho hàm json.Marshal và nó sẽ mã hóa dữ liệu như bạn mong đợi:
func MarshalingSlices() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// Now we pass a slice of two pigeons
	data, _ := json.Marshal([]*Bird{pigeon, pigeon})
	fmt.Println(string(data))
}

// Kiểu gì cũng có thể convert được qua json hết và ngược lại
func MarshalingMaps() {
	birdData := map[string]any{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squawk",
		},
		"total birds": 2,
	}

	// JSON encoding is done the same way as before
	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}

type Bird5 struct {
	Species     string
	Description *string
}

// Muốn sử dụng giá trị null để truyền qua json thì buộc filed nhận giá trị null phải được định nghĩa là kiểu con trỏ
func EncodingNullValues() {
	pigeon := &Bird5{
		Species:     "Pigeon",
		Description: nil,
	}

	// way 1:
	data, _ := json.Marshal(pigeon)
	fmt.Println(string(data))

	// way 2:
	birdData := map[string]any{
		"total birds": 2,
		"ok":          nil,
	}

	data, _ = json.Marshal(birdData)
	fmt.Println(string(data))
}

type Dimensions2 struct {
	Height int
	Width  int
}

type Bird6 struct {
	Species    string
	Dimensions Dimensions2
}

// Đây là một tính năng mạnh mẽ trong Go cho phép bạn kiểm soát chính xác cách các cấu trúc dữ liệu của bạn được chuyển đổi thành JSON
// Hàm này sẽ tự động được gọi khi một đối tượng Dimensions2 được mã hóa thành JSON.

func (d Dimensions2) MarshalJSON() ([]byte, error) {
	fmt.Print("AUTO MarshalJSON\n")
	return []byte(fmt.Sprintf(`"%dx%d"`, d.Height, d.Width)), nil
}

func CustomEncodingLogic() {
	bird := Bird6{
		Species: "pigeon",
		Dimensions: Dimensions2{
			Height: 24,
			Width:  10,
		},
	}
	birdJson, _ := json.Marshal(bird)
	fmt.Println(string(birdJson))
}

// Printing Formatted (Pretty-Printed) JSON (beautiful print code)
func PrintingFormattedPrettyPrintedJSON() {
	bird := Bird{
		Species:     "pigeon",
		Description: "likes to eat seed",
	}

	// The second parameter is the prefix of each line, and the third parameter
	// is the indentation to use for each level
	data, _ := json.MarshalIndent(bird, "", "  ")
	fmt.Println(string(data))

}

// Practice homework: write system about convert restApi to Struct dynamic and otherwise

type JSON struct {
	Type interface{}
	Data string
}

type Input struct {
	data interface{}
}

type Output struct {
	data interface{}
}

type User struct {
	Name string `json:"ten"`
	Age  int    `json:"tuoi"`
}

func (j *JSON) Stringify() interface{} {
	// Convert struct to string

	json.Unmarshal([]byte(j.Data), &j.Type)

	return j.Type
}

func (j *JSON) Parse() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "createdAt": "2021-10-18T11:08:47.577Z"}`
	var bird Bird2
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)

}

func PracticeJsonDynamic() {
	var (
		body = `{"ten":"123", "tuoi": "312"}`
	)

	var data User

	json.Unmarshal([]byte(body), &data)

	// var convert = JSON{Type: User{}, Data: body}

	fmt.Println(data)
}
