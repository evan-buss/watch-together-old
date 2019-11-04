// import { writable } from "svelte/store";

export let socket = null;
export function createSocket(){
  if (socket === null) {
    socket = new SocketHandler("ws://localhost:8080/ws");
  }
}

// SocketHandler wraps the websocket and 
// adds an event handler that parses the event type and data from a JSON message
export function SocketHandler(url) {
  let conn = new WebSocket(url);

  let callbacks = {}

  this.bind = function (name, callback) {
    callbacks[name] = callbacks[name] || [];
    callbacks[name].push(callback);
    return this;
  }

  this.send = function (name, data) {
    let payload = JSON.stringify({ event: name, data: data })
    conn.send(payload);
    return this;
  }

  this.close = () => {
    conn.close();
  }

  conn.onmessage = event => {
    let json = JSON.parse(event.data);
    dispatch(json.event, json.data);
  }

  conn.onclose = () => {
    dispatch('close', null)
  }

  conn.onopen = () => {
    dispatch('open', null)
  }

  let dispatch = function (event, message) {
    let chain = callbacks[event];
    if (typeof chain == 'undefined') return;
    for (var i = 0; i < chain.length; i++) {
      chain[i](message)
    }
  }

}