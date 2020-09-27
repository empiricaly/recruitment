export const debounce = (func, delay) => {
  let inDebounce;
  return function () {
    const context = this;
    const args = arguments;
    clearTimeout(inDebounce);
    inDebounce = setTimeout(() => func.apply(context, args), delay);
  };
};

export const throttle = (fn, threshhold) => {
  threshhold || (threshhold = 250);
  var last, deferTimer;
  return function () {
    var now = +new Date(),
      args = arguments;
    if (last && now < last + threshhold) {
      // hold on to it
      clearTimeout(deferTimer);
      deferTimer = setTimeout(function () {
        last = now;
        fn.apply(this, args);
      }, threshhold);
    } else {
      last = now;
      fn.apply(this, args);
    }
  };
};
