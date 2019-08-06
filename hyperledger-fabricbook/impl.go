package chaincode


import (
	"cartransfer"
	"encoding/json"
	"errors"
	"fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func checkLen(logger *shim.ChaincodeLogger, expected int, args []string) error{
	if len(args) < expected {
		mes := fmt.Sprintf(
			"not enough number of arguments: %d given, %d expected",
			len(args),
			expected,
		)
		logger.Warning(mes)
		return errors.New(mes)
	}
	return nil
}


type CarTransferCC struct {
}

func (this *CarTransferCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger := shim.NewLogger("cartransfer")
	logger.Info("chaincode initialized")
	return shim.Success([]byte{})
}

func (this *CarTransferCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger := shim.NewLogger("cartransfer")

	//sample of API use: show tX timestamp
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get TX timestamp: %s", err))
	}
	logger.Infof(
		"Invoke called: Tx ID = %s, timestamp = %s",
		stub.GetTxID(),
		timestamp,
	)
	var (
		fcn string
		args []string
	)
	fcn, args = stub.GetFunctionAndParameters()
	logger.Infof("function name = %s", fcn)

	switch fcn {
		// adds a new Owner
	case "AddOwner":
		// checks arguments length
		if err := checkLen(logger, 1, args)1 err != nil {
			return shim.Error(err.Error())
		}

		// unmarshal
		owner := new(cartransfer.Owner)
		err := json.Unmarshal([]byte(args[0]), owner)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal Owner JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.AddOwner(stub, owner)
		if err != nil {
			returnsshim.Error(err.Error())
		}

		// returns a success value
		return shim.Success([]byte{})
	}

	// lists Owners
case "ListOwners":
	owners, err := this.ListOwners(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	// marshal
	b, err := json.Marshal(owners)
	if err != nil {
		mes := fmt.Sprintf("failed to marshal Owners: %s", err.Error())

		logger.Warning(mes)
		return shim.Error(mes)
	}
	// returns a success value
	return shim.Success(b)
	// adds a new Car
case "AddCar":
	// checks arguments length
	if err := checkLen(logger, 1, args); err != nil {
		return shim.Error(err.Error())
	}

	// unmarshal
	car := new(cartransfer.Car)
	err := json.Unmarshal([]byte(args[0]), car)
	if err != nil {
		mes := fmt.Sprintf("failed to unmarshal Car JSON: %s", err.Error())
		logger.Warning(mes)
		return shim.Error(mes)
	}

	err = this.AddCar(stub, car)
	if err != nil {
		return shim.Error(err.Error())
	}

	// returns a success value
	return shim.Success([]byte{})

	// lists Cars
case "ListCars":
	cars, err := this.ListCars(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	// marshal
	b, err := json.Marshal(cars)
	if err != nil {
		mes := fmt.Sprintf("failed to marshal Cars: %s", err.Error())
		logger.Warning(mes)
		return shim.Error(mes)
	}

	// returns a success value
	return shim.Success(b)

	// gets an existing Car

case "GetCar":
	// checks arguments length
	if err := checkLen(logger, 1, args); err != nil {
		return shim.Error(err.Error())
	}

	// unmarshal
	var id string
	err := json.Unmarshal([]byte(args[0]), &id)
	if err != nil {
		mes := fmt.Sprintf("failed to unmarshal the 1st argument: %s", err.Error())
		logger.Warning(mes)
		return shim.Error(mes)
	}

	car, err := this.GetCar(stub, id)
	if err != nil {
		return shim.Error(err.Error())
	}

	// marshal
	b, err := json.Marshal(car)
	if err != nil {
		mes := fmt.Sprintf("failed to marshal Car: %s", err.Error())
		logger.Warning(mes)
		return shim.Error(mes)
	}

	// returns a success value
	return shim.Success(b)

	// updates an existing Car
case "UpdateCar":
	// checks arguments length
	if err := checkLen(logger, 1, args); err != nil {
		return shim.Error(err.Error())
	}

	// unmarshal
	car := new(cartransfer.Car)
	err := json.Unmarshal([]byte(args[0]), car)
	if err != nil {
		mes := fmt.Sprintf("failed to unmarshal Car JSON: %s", err.Error())
		logger.Warning(mes)
		return shim.Error(mes)
	}

	err = this.UpdateCar(stub, car)
	if err != nil {
		return shim.Error(err.Error())
	}

	// returns a success value
	return shim.Success([]byte{})

	// transfers an existing Car to an existing Owner
case "transferCar":
	// checks arguments length
	if err := checkLen(logger, 2, args); err != nil {
		return shim.Error(err.Error())
	}

	// unmarshal
	var carId, newOwnerId string
	err := json.Unmarshal([]byte(args[0]), &carId)
	if err != nil {
		mes := fmt.Sprintf(
			"failed to unmarshal the 1st argument: %s",
			err.Error(),
		)
		logger.Warning(mes)
		return shim.Error(mes)
	}

	err = json.Unmarshal([]byte(args[1]), &newOwnerId)
	if err != nil {
		mes := fmt.Sprintf(
			"failed to unmarshal the 2nd argument: %s",
			err.Error(),
		)
		logger.Warning(mes)
		return shim.Error(mes)
	}

	err = this.TransferCar(stub, carId, newOwnerId)
	if err != nil {
		return shim.Error(err.Error())
	}

	// returns a success value
	return shim.Success([]byte{})
	
	// if the function name is unknown
	mes := fmt.Sprintf("Unknown method: %s", fcn)
	logger.Warning(mes)
	return shim.Error(mes)
}

// methos implementing CarTransfer interface

// Adds a new Owner
func (this *CarTransferCC) AddOwner(stub shim.ChaincodeStubInterface, owner *cartransfer.Owner) error{
	return errors.New("not implemented yet")
}

// checks existence of the specified Owner
func (this *CarTransferCC) CheckOwner(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	return false, errors.New("not implemented yet")
}

// Lists Owners
func (this *CarTransferCC) ListOwners(stub shim.ChaincodeStubInterface) ([]*cartransfer.Owner, error) {
	return nil, errors.New("not implemented yet")
}

// Adds a new Car
func (this *CarTransferCC) AddCar(stub shim.ChaincodeStubInterface, car *cartransfer.Car) error {
	return errors.New("not implemented yet")
}

// checks existence of the specified Car
func (this *CarTransferCC) CheckCar(stub shim.ChaincodeStubInterface, id string) (bool, error){
	return false, errors.New("not implemented yet")
}

// Validates the content of the specified Car
func (this *CarTransferCC) ValidateCar(stub shim.ChaincodeStubInterface, car *cartransfer.Car) (bool, error) {
	return false, errors.New("not implemented yet")
}

// Gets the specified Car
func (this *CarTransferCC) GetCar(stub shim.ChaincodeStubInterface, id string) (*cartransfer.Car, error) {
	return nil, errors.New("not implemented yet")
}

// Updates the content of the specified Car
func (this *CarTransferCC) UpdateCar(stub shim.ChaincodeStubInterface, car *cartransfer.Car) error {
	return errors.New("not implemented yet")
}

// Lists cars
func (this *CarTransferCC) ListCars(stub shim.ChaincodeStubInterface) ([]*cartransfer.Car, error) {
	return nil, errors.New("not implemented yet")
}

// Transfers the specified Car to the specified Owner
func (this *CarTransferCC) TransferCar(stub shim.ChaincodeStubInterface, carId string, newOwnerId string) error{
	return nil, errors.New("not implemented yet")
}