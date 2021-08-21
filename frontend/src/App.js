import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

import Login from "./components/login";

import SignUp from "./components/signup";

function App() {
  return (
    <Router>
      <div>
       <Switch>
          <Route exact path="/">
            <SignUp />
          </Route>
          <Route exact path="/login">
            <Login />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
