function Store() {
    console.log(`${this.constructor.name}-Store:constructor:enter+leave`)
    this.observers = []; 
}

Store.prototype.registerObserver = function(observer) {
    this.observers.push(observer);
}

Store.prototype.notifyAll = function() {
    this.observers.forEach(function(observer) { observer(this)}.bind(this))
}


function PositionStore() {
    Store.call(this)
    this.pos = null
}

Object.setPrototypeOf(PositionStore.prototype, Store.prototype);

PositionStore.prototype.getPos = function() {
    return this.pos;
}

PositionStore.prototype.setPos = function(pos) {
    this.pos = pos
    this.notifyAll()
}

/////////////////////////// SpotStore

function SpotStore(args) {
    // call parent constructor
    Store.call(this)
    this.spot = null;
}

// SpotStore is child of Store
Object.setPrototypeOf(SpotStore.prototype, Store.prototype);

SpotStore.prototype.getSpot= function() {
    return this.spot;
}

SpotStore.prototype.setSpot = function(spot) {
    this.spot = spot;
    this.notifyAll()
}
