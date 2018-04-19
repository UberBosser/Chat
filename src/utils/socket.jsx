import {EventEmitter} from "events";


class Socket {
    constructor(ws = new WebSocket("ws://localhost:8080/socket"), ee = new EventEmitter()) {
        this.ws = ws;
        this.ee = ee;
        ws.onopen = this.open.bind(this);
        ws.onclose = this.close.bind(this);
        ws.onmessage = this.message.bind(this);
    }
    on(code, fn) {
        this.ee.on(code, fn); 
    }
    off(code, fn) {
        this.ee.removeListener(code, fn);
    }
    emit(code, data) {
        const message = JSON.stringify({code, data});
        this.ws.send(message);
    }
    message(e) {
        try {
            const message = JSON.parse(e.data);
            this.ee.emit(message.code, message.data);
        }
        catch(err) {
            this.ee.emit("error", err);
        }
    }
    open() {
        this.ee.emit("open");
    }
    close() {
        this.ee.emit("close");
    }
}


export default Socket;
