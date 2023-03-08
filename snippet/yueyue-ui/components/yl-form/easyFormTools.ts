export function easyFormGetProperty(raw: AnyObject, name: string): AnyObject {
  const nameArr = name.split(".");
  if (nameArr.length >= 2) {
    let obj = { ...raw };
    nameArr.forEach((item) => {
      obj = obj[item];
    });
    return obj;
  }
  return raw[name];
}

export function easyFormSetProperty(
  raw: AnyObject,
  name: string,
  value: AnyObject
) {
  const nameArr = name.split(".");
  if (nameArr.length >= 2) {
    let obj = raw[nameArr[0]];
    for (let i = 1; i < nameArr.length - 1; i++) {
      obj = obj[nameArr[i]];
    }
    obj[nameArr[nameArr.length - 1]] = value;
  } else {
    raw[name] = value;
  }
}
