import "./App.css";
import { BrowserRouter as Router, Route } from "react-router-dom";
import ShortenURL from "./components/createShortURL/shorturl/";
import Navigate from "./components/navigate";

function App() {
  return (
    <div className="App">
      <Router>
        <Route path="/:url">
          <Navigate />
        </Route>
        <Route path="/" exact>
          <ShortenURL />
        </Route>
      </Router>
    </div>
  );
}

export default App;
