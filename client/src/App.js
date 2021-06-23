import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Login from './pages/Login';
import Register from './pages/Register';
import styled from 'styled-components';
import Home from './pages/Home';
import AddPassword from './pages/AddPassword';
import EditPassword from './pages/EditPassword';
import Profile from './pages/Profile';
import UpdateProfile from './pages/ProfileUpdate';
import Dashboard from './pages/Dashboard';

function App() {
  const App = styled.div`
  position: relative;
  min-height: 100vh;
  `;

  return (
    <App className="App">
      <Router>
        <Switch>
          <Route path="/register" exact component={Register} />
          <Route path="/login" exact component={Login} />
          <Route path="/password/add" exact component={AddPassword} />
          <Route path="/password/edit/:id" exact component={EditPassword} />
          <Route path="/profile/:id" exact component={Profile} />
          <Route path="/profile/edit/:id" exact component={UpdateProfile} />
          <Route path="/dashboard" exact component={Dashboard} />
          <Route path="/" exact component={Home} />
        </Switch>
      </Router>
    </App>
  );
}

export default App;
