import {Message} from 'view-design';

/**
 * showMessage 全局显示提示消息
 * @params {*} type
 * @params {*} content
 */
export const showMessage = (type, content) => {
    return Message[type] && Message[type]({
        content,
        duration: 3,
        closable: true
    });
};

/**
 * setLocalStorage 设置localStorage
 * @params {*} name
 * @params {*} value
 */
export const setLocalStorage = (name, value) => {
    window.localStorage.setItem(name, value);
    console.log('设置成功');
}

/**
 * getLocalStorage
 * @params {*} name
 */
export const getLocalStorage = (name) => {
    console.log('获取localStroge');
    return window.localStorage.getItem(name);
}

/**
 * removeLocalStorage
 * @params {*} name
 */
export const removeLocalStorage = (name) => {
    return window.localStorage.removeItem(name);
}
