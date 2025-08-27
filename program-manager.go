package bounded_xdg

import (
	"fmt"
)

func NewProgramManager(fm *BoundedFileManager) *BoundedProgramManager {
	return &BoundedProgramManager{
		availableRoot: "$BOUNDED_CONFIG/programs/available",
		enabledRoot:   "$BOUNDED_CONFIG/programs/enabled",
		embeddedRoot:  "programs",
		fm:            fm,
	}
}

type BoundedProgramManager struct {
	availableRoot string
	enabledRoot   string
	embeddedRoot  string
	fm            *BoundedFileManager
	stock         map[string]BoundedProgramConfig
	available     map[string]BoundedProgramConfig
	enabled       map[string]bool
}

func (this *BoundedProgramManager) init(config BoundedProgramConfig) error {
	embeddedFiles, err := this.fm.embeddedReadDir(this.embeddedRoot)

	if err != nil {
		return err
	}

	for _, file := range embeddedFiles {
		var config BoundedProgramConfig
		if err := this.fm.embeddedUnmarshalFile(&config, this.embeddedRoot, file.Name()); err != nil {
			return err
		}
		this.stock[config.Name] = config
	}

	availableFiles, err := this.fm.fsReadDir(this.availableRoot)

	if err != nil {
		return err
	}

	for _, file := range availableFiles {
		var config BoundedProgramConfig
		if err := this.fm.fsUnmarshalFile(&config, this.availableRoot, file.Name()); err != nil {
			return err
		}
		this.available[file.Name()] = config
	}

	enabledFiles, err := this.fm.fsReadDir(this.enabledRoot)

	if err != nil {
		return err
	}

	for _, file := range enabledFiles {
		this.enabled[file.Name()] = true
	}

	return nil
}

func (this *BoundedProgramManager) Available(config BoundedProgramConfig) error {
	filename := fmt.Sprintf("%s.json", config.Name)
	exists, err := this.fm.fsExists(this.availableRoot, filename)

	if err != nil {
		return err
	}

	if exists.Exists {

	}

	return nil
}
