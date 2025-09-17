components {
  id: "screen_proxy"
  component: "/monarch/screen_proxy.script"
  properties {
    id: "screen_id"
    value: "reload_game_screen"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "popup"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "timestep_below_popup"
    value: "0.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "collectionproxy"
  type: "collectionproxy"
  data: "collection: \"/screens/reload_game_screen/reload_game_screen.collection\"\n"
}
