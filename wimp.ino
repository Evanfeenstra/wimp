#include <MIDI.h>

#define USE_RUNNING_STATUS 0

//MIDI_CREATE_INSTANCE(HardwareSerial, Serial1, midi1); //Serial1 or SerialUSB
MIDI_CREATE_INSTANCE(HardwareSerial, Serial2, midi2);
MIDI_CREATE_INSTANCE(HardwareSerial, Serial3, midi3);

int channel = 1;

void setup(){
  pinMode(LED_BUILTIN, OUTPUT);
  //digitalWrite(LED_BUILTIN, HIGH);
  Serial.begin(115200);
  midi2.begin(MIDI_CHANNEL_OMNI);
  midi3.begin(MIDI_CHANNEL_OMNI);
  while(!Serial2);
  while(!Serial3);
  midi2.turnThruOff();
  midi3.turnThruOff();

  midi2.setHandleNoteOn(Midi2NoteOn);
  midi2.setHandleNoteOff(Midi2NoteOff);
  midi3.setHandleNoteOn(Midi3NoteOn);
  midi3.setHandleNoteOff(Midi3NoteOff);

  midi2.setHandleControlChange(Midi2ControlChange);
  midi3.setHandleControlChange(Midi3ControlChange);
}

elapsedMillis flash;
void loop(){
  midi2.read();
  midi3.read();
//  int note = random(40,100);
//  Serial.print("1:"+String(note)+":100");  
//  delay(1000);
  if (flash>100){
    flash = 0;
    digitalWrite(LED_BUILTIN, LOW);
  }
  
}

void Midi2NoteOn(byte channel, byte pitch, byte velocity) {
  Serial.print("1:"+String(pitch)+":100");
  digitalWrite(LED_BUILTIN, HIGH);
  flash = 0;
}
void Midi2NoteOff(byte channel, byte pitch, byte velocity) {
  //Serial.println("midi2noteon");
}
void Midi3NoteOn(byte channel, byte pitch, byte velocity) {
  Serial.print("1:"+String(pitch)+":100");
  digitalWrite(LED_BUILTIN, HIGH);
  flash = 0;
}
void Midi3NoteOff(byte channel, byte pitch, byte velocity) {
  //Serial.println("midi3noteon");
}
void Midi2ControlChange(byte channel, byte number, byte val){   
  Serial.print("c:"+String(number)+":"+String(val));
}
void Midi3ControlChange(byte channel, byte number, byte val){   
  Serial.print("C:"+String(number)+":"+String(val));
}
