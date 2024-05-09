import Sdk from "casdoor-js-sdk";

export const ServerUrl = "http://localhost:8080";

const sdkConfig = {
    serverUrl: "https://door.casdoor.com",
    clientId: "294b09fbc17f95daf2fe",
    appName: "app-vue-python-example",
    organizationName: "casbin",
    redirectPath: "/callback",
};
  
export const CasdoorSDK = new Sdk(sdkConfig);

export const isLoggedIn = () => {
    const token = localStorage.getItem("token");
    return token !== null && token.length > 0;
};
  
export const setToken = (token) => {
    localStorage.setItem("token", token);
};
  
export const goToLink = (link) => {
    window.location.href = link;
};
  
export const getSigninUrl = () => {
    return CasdoorSDK.getSigninUrl();
};

export const logout = () => {
    localStorage.removeItem("token");
};

export const showMessage = (message) => {
    alert(message);
};

