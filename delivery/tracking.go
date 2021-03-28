package delivery

import "fmt"

func (d *deliver) GetTracking() {
	result, err := d.usecase.GetTracking()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Print(result)
}
