function Errors() {

    this.errors = [];

    this.observers = []; 
}

Identity.prototype.registerObserver = function(observer) {
    this.observers.push(observer);
}

Identity.prototype.notifyAll = function() {
    this.observers.forEach(function(observer) { observer(this)}.bind(this))
}