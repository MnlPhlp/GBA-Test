tinygo build -target gameboy-advance -o tmp test.go
cp tmp MnlPhlp_emu.gba
arm-none-eabi-objcopy -O binary tmp MnlPhlp.gba
rm tmp
gbafix MnlPhlp.gba -tMnlPhlp