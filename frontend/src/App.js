import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

import SignUp from "./components/signup";

function App() {
  return (
    <Router>
      <div>
       <Switch>
          <Route path="/">
            <SignUp />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
