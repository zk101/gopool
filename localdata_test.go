package gopool

import "errors"

// A LocalData implementation for testing purposes

func testLocalDataFactory() LocalData {
	return &testLocalData{}
}

func testLocalDataFactoryWithError() LocalData {
	return &testLocalData{
		forceError: true,
	}
}

type testLocalData struct {
	forceError bool
}

func (tld *testLocalData) Setup() error {
	if tld.forceError {
		return errors.New("forcing a localdata setup error")
	}

	return nil
}

// EOF
