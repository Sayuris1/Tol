components {
  id: "robotController"
  component: "/scrpits/robotController.script"
  properties {
    id: "speed"
    value: "500.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "circle_factory"
  type: "factory"
  data: "prototype: \"/prefabs/circle.go\"\n"
  ""
}
embedded_components {
  id: "weapon_factory"
  type: "factory"
  data: "prototype: \"/prefabs/weapon.go\"\n"
  ""
}
