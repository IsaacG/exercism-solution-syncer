// @ts-check

/**
 * Implement the classes etc. that are needed to solve the
 * exercise in this file. Do not forget to export the entities
 * you defined so they are available for the tests.
 */

export function Size(width = 80, height = 60) {
  this.width = width
  this.height = height
}

Size.prototype.resize = function (width, height) {
  this.width = width
  this.height = height
}

export function Position(x = 0, y = 0) {
  this.x = x
  this.y = y
}

Position.prototype.move = function (newX, newY) {
  this.x = newX
  this.y = newY
}

export class ProgramWindow {
  constructor() {
    this.screenSize = new Size(800, 600)
    this.size = new Size()
    this.position = new Position()
  }

  resize(size) {
    let w = size.width > 0 ? size.width : 1;
    let h = size.height > 0 ? size.height : 1;
    const maxW = this.screenSize.width - this.position.x
    const maxH = this.screenSize.height - this.position.y
    if (w > maxW) { w = maxW }
    if (h > maxH) { h = maxH }
    this.size.resize(w, h)
  }

  move(position) {
    let x = position.x >= 0 ? position.x : 0
    let y = position.y >= 0 ? position.y : 0
    const maxX = this.screenSize.width - this.size.width
    const maxY = this.screenSize.height - this.size.height
    if (x > maxX) { x = maxX }
    if (y > maxY) { y = maxY }
    this.position.move(x, y)
  }
}

export function changeWindow(window) {
  window.size.resize(400, 300)
  window.position.move(100, 150)
  return window
}
