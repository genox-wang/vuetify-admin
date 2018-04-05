const prefix = 'vuetify_admin_';

class Store {
  constructor() {
    this.store = window.localStorage;
    this.prefix = prefix;
  }
  set(key, value) {
    let v = value;
    try {
      v = JSON.stringify(value);
    } catch (e) {
      throw new Error(e);
    }

    this.store.setItem(this.prefix + key, v);
  }
  get(key) {
    if (!key) {
      throw new Error('没有找到key。');
    }
    if (typeof key === 'object') {
      throw new Error('key不能是一个对象。');
    }
    let value = this.store.getItem(this.prefix + key);
    if (value !== null) {
      try {
        value = JSON.parse(value);
      } catch (e) {
        throw new Error(e);
      }
    }

    return value;
  }
  remove(key) {
    this.store.removeItem(this.prefix + key);
  }
}

export default new Store();
