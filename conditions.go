package iso8583

func (m *Message) GetMandatoryFields() (mandatory, echo, ifAvailable []int, err error) {
	flow, _, err := m.GetFlow()
	if err != nil {
		return
	}

	for index, requirement := range flow.MandatoryFields {

		switch requirement {
		case "O":
			continue
		case "M", "03":
			mandatory = append(mandatory, index)
		case "ME":
			echo = append(echo, index)
		case "01", "02":
		case "04":
		case "05":
		case "06":
		case "07":
		case "08":
		case "09":
		case "10":
		case "11":
		case "12":
		case "13":
		case "14":
		case "15":
		case "16":
		case "17":
		case "18":
		case "19":
		case "20":
		case "21":
		case "22":
		case "23":
		case "24":
		case "25":
		case "26":
		case "27":
		case "28":
		case "29":
		case "30":
		case "31":
		case "32":
		case "33":
		case "34":
		case "35":
		case "36":
		case "37":
		case "38":
		case "39":
		case "40":
		case "41":
		case "42":
		case "43":
		case "44":
		case "45":

		}
	}

	return
}
