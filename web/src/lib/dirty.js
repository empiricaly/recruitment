let dirtyObject = {};

export const addDirtyObject = (id) => (dirtyObject[id] = true);

export const removeDirtyObject = (id) => delete dirtyObject[id];

window.onbeforeunload = function (e) {
  if (Object.keys(dirtyObject).length > 0) {
    e.preventDefault();
    e.returnValue = "";
  }
};
