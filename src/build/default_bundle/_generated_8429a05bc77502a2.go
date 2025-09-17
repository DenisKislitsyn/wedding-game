components {
  id: "screen_proxy"
  component: "/monarch/screen_proxy.script"
  properties {
    id: "screen_id"
    value: "win_screen"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "popup"
    value: "false"
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
  data: "collection: \"/screens/win_screen/win_screen.collection\"\n"
}
