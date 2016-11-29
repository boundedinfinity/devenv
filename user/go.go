package user

func (this *UserConfigManager) EnableGo() error {
    if err := this.EnsureConfigDir(); err != nil {
        return err
    }



    return nil
}

func (this *UserConfigManager) DisableGo() error {
    if err := this.EnsureConfigDir(); err != nil {
        return err
    }

    return nil
}
