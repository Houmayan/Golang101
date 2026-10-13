package train

import "fmt"

func GetTrainScheduleUpdate(typeOfTrain string, delayTime int) string {

	switch typeOfTrain {
	case "ICE":
		if delayTime  > 0 {
			return fmt.Sprintf("%s Trian is delayed by %d minute",typeOfTrain,delayTime)
		}
				return "✅ The ICE is arriving on time on Platform 1."
	case "RE":
		if delayTime > 0 {
			return  fmt.Sprintf("%s Delayed by %d minutes ",typeOfTrain,delayTime)
		}
				return "✅ The ICE is arriving on time on Platform 2."
	case "S-Bahn":
		if delayTime  > 0 {
			return  fmt.Sprintf("%s Delayed by %d minutes ",typeOfTrain,delayTime)
		}
				return "✅ The ICE is arriving on time on Platform 3."
	case "Bus":
		if delayTime  > 0 {
			return  fmt.Sprintf("%s Delayed by %d minutes ",typeOfTrain,delayTime)
		}
				return "✅ The ICE is arriving on time on Platform 4."
	case "Euro City":
		if delayTime  > 0 {
			return  fmt.Sprintf("%s Delayed by %d minutes ",typeOfTrain,delayTime)
		}
				return "✅ The ICE is arriving on time on Platform 5."
	default:
		return  fmt.Sprintf("Please Check announcement Board")
	}

}
