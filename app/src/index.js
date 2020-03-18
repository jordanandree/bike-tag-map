import React from 'react';
import ReactDOM from 'react-dom';
import Post from './Post';

const apiUrl = 'http://localhost:8080';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
      error: null,
    };
  }

  componentDidMount() {
    fetch(`${apiUrl}/posts`)
      .then(res => res.json())
      .then(
        result => {
          this.setState({
            items: result,
          });
        },
        error => {
          this.setState({
            error,
          });
        }
      );
  }

  render() {
    const { items, error } = this.state;

    if (error) {
      return (
        <div>
          Uh oh! An error occured:
          <br />
          {error.message}
        </div>
      );
    }

    return (
      <div>
        {items.map(item => {
          return <Post key={item.id} id={item.id} title={item.title} />;
        })}
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
