import * as util from "./util";
import { SilentSignin, isSilentSigninRequired } from "casdoor-react-sdk";

function HomePage() {
  const casdoorLogin = () => {
    util.goToLink(util.CasdoorSDK.getSigninUrl());
  }

  if (isSilentSigninRequired()) {
    return (
      <SilentSignin
        sdk={util.CasdoorSDK}
        isLoggedIn={util.isLoggedIn}
        handleReceivedSilentSigninSuccessEvent={() => {
          util.goToLink("/");
        }}
        handleReceivedSilentSigninFailureEvent={() => {
          alert("login failed");
        }}
      />
    );
  }

  if (util.isLoggedIn()) {
    return <div>Hello!</div>;
  }

  return  (
    <div
        style={{
          textAlign: "center",
          alignItems: "center",
        }}
      >
        <button onClick={casdoorLogin}>Casdoor Login</button>
      </div>
  );
}

export default HomePage;