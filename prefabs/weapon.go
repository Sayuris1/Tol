components {
  id: "weaponController"
  component: "/scrpits/weaponController.script"
  properties {
    id: "fire_interval"
    value: "0.5"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "bullet_speed"
    value: "1000.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/builtins/graphics/particle_blob.tilesource\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "bullet_factory"
  type: "factory"
  data: "prototype: \"/prefabs/bullet.go\"\n"
  ""
}
