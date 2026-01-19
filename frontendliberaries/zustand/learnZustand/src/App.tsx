

import './App.css'

function App() {


  return (
    <>
      <OtherComponent userName="Dhanshree Upadhye" />
    </>
  );
}

export default App


type OtherComponentProps = {
  userName: string;
}

const OtherComponent : React.FC<OtherComponentProps> = ({userName}) => {

  return <div>
<h1>other component {userName}</h1>
<p>This is another component.</p>


  </div>
}

//export default otherComponent;