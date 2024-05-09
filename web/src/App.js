import "./App.css";

import { Route, BrowserRouter, Routes } from "react-router-dom";
import { AuthCallback } from "casdoor-react-sdk";
import * as util from "./util";
import HomePage from "./HomePage";

function App() {
  const authCallback = (
    <AuthCallback
      sdk={util.CasdoorSDK}
      serverUrl={util.ServerUrl}
      saveTokenFromResponse={(res) => {
        util.setToken(res.data);
        util.goToLink("/");
      }}
      isGetTokenSuccessful={(res) => {
        console.log(res.status);
        return res.status === "ok";
      }}
    />
  );

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/callback" element={authCallback} />
        <Route path="/" element={<HomePage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;