import React from "react"
import './App.css'

import RegisterPage from "./pages/register"
import LoginPage from "./pages/login";
import TablePass from "./pages/passMan";
import CreatePass from "./pages/passMancreate";
import UpdatePass from "./pages/updatePassMan";

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"; 

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/register">
          <RegisterPage/>
        </Route>
        <Route path="/login">
          <LoginPage/>
        </Route>
        <Route path="/site">
          <TablePass/>
        </Route>
        <Route path="/add-site">
          <CreatePass/>
        </Route>
        <Route path="/edit-site/:pass_id">
          <UpdatePass/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
