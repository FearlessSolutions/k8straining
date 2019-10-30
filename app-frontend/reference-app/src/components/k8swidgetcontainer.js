import React, { useState, useEffect }  from 'react';
import Accordion from "react-bootstrap/Accordion";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";

export default function K8WidgetContainer(props) {

    const [results, setResults] = useState([]);
    const [nestActive, setNestActive] = useState(false);
    const [nestInterval, setNestInterval] = useState(13);
    const [addA, setAddA] = useState(3);
    const [addB, setAddB] = useState(7);
    const [multiplyA, setMultiplyA] = useState(5);
    const [multiplyB, setMultiplyB] = useState(15);
    const [url, setUrl] = useState(
        '',
    );

    useEffect(() => {
        const fetchData = async () => {
            if (url !== '') {
                let newResults = results.slice();
                newResults.unshift(`REQUESTING FROM:${url}`);
                let fr = await fetch("https://hn.algolia.com/api/v1/search?query=redux");
                newResults.unshift(`RESULT is: ${addA + addB}`);
                setResults(newResults);
            }
        };
        fetchData();
    }, [url]);

    function reqAdd() {
        let desiredUrl = `http://localhost:3000/add/?a=${addA}&b=${addB}`;
        setUrl(desiredUrl);
    }

    function reqMultiply() {
        let desiredUrl = `http://localhost:3000/multiply/?a=${multiplyA}&b=${multiplyB}`;
        setUrl(desiredUrl);
    }

    return (
        <div className='container'>
            <div className="row">
                <div className={"col"}>
                    <h2>ENDPOINTS</h2>
                    <Card>
                        <Card.Header>
                                Nest

                        </Card.Header>
                        <Card.Body>
                            Send requests to the nest endpoint every
                            <input className={"col-sm-2"} value={nestInterval} onChange={(e) =>{setNestInterval(e.target.value)}}/> microseconds
                            <br/>
                            NEST interval is: {nestInterval}
                            <br/>
                            {nestActive ?
                                <Button onClick={() => setNestActive(!nestActive)} variant="primary" size="md" active>
                                    Deactivate Nest
                                </Button>
                                :
                                <Button onClick={() => setNestActive(!nestActive)} variant="primary" size="md">
                                    Activate Nest
                                </Button>
                            }
                        </Card.Body>
                    </Card>
                    <Card>
                        <Card.Header>
                            Add
                        </Card.Header>
                        <Card.Body>
                            send request to the add endpoint to add:
                            <br/>
                            <input className={"col-sm-2"} value={addA} onChange={(e) =>{setAddA(e.target.value)}}/>
                            to
                            <input className={"col-sm-2"} value={addB} onChange={(e) =>{setAddB(e.target.value)}}/>
                            <br/>
                            <Button onClick={()=>reqAdd()}>Send</Button>
                        </Card.Body>
                    </Card>
                    <Card>
                        <Card.Header>
                            Multiply
                        </Card.Header>
                        <Card.Body>
                            send request to the multiply endpoint to multiply:
                            <br/>
                            <input className={"col-sm-2"} value={multiplyA} onChange={(e) =>{setMultiplyA(e.target.value)}}/>
                            by
                            <input className={"col-sm-2"} value={multiplyB} onChange={(e) =>{setMultiplyB(e.target.value)}}/>
                            <br/>
                            <Button onClick={()=>reqMultiply()}>Send</Button>
                        </Card.Body>
                    </Card>
                </div>
                <div className="col">
                    <h2>RESULTS</h2>
                    <ul>
                        {results.map((result) => <li>{result}</li>)}
                    </ul>
                </div>
            </div>
        </div>
    )
}

