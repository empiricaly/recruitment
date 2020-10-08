export const debounce = (func, delay, max = 0) => {
  let inDebounce;
  let first;
  return function () {
    const context = this;
    const args = arguments;
    clearTimeout(inDebounce);
    let wait = delay;
    if (max > 0) {
      if (!first) {
        first = new Date();
      } else {
        const ellapsed = new Date() - first;
        wait = Math.min(max - ellapsed, delay);
      }
    }
    inDebounce = setTimeout(() => {
      first = null;
      func.apply(context, args);
    }, wait);
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
