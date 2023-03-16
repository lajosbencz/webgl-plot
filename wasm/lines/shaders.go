package main

const vsSource = `
attribute vec2 aPosition;
void main(void) {
  gl_Position = vec4(aPosition, 0.0, 1.0);
}
`

const fsSource = `
void main(void) {
  gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
}
`

const vsSourceColor = `
precision highp float;
attribute vec2 aPosition;
attribute vec3 aColor;
varying vec3 vColor;
void main(void) {
  gl_Position = vec4(aPosition, 0.0, 1.0);
  vColor = aColor;
}
`

const fsSourceColor = `
precision highp float;
varying vec3 vColor;
void main(void) {
  gl_FragColor = vec4(vColor, 1.0);
}
`
