tinygo build -target gameboy-advance -o tmp test.go
cp tmp game_emu.gba
arm-none-eabi-objcopy -O binary tmp game.gba
rm tmp
gbafix game.gba -tgame
