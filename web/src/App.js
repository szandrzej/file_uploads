import React, { useState, useEffect, useRef } from 'react';
import axios from 'axios';

const API_URL = '/api/files/'
const OPTIONS = { headers: { 'Authorization': 'example_token'}}

const App = () => {
  const [file, setFile] = useState(null);
  const [filesList, setFilesList] = useState([]);
  const fileInputRef = useRef();

  // Fetch the list of files from the server
  useEffect(() => {
    axios.get(API_URL, OPTIONS)
      .then(response => setFilesList(response.data.files))
      .catch(error => console.error('Error fetching files:', error));
  }, []);

  // Handle file input change
  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  // Handle file upload
  const handleFileUpload = () => {
    const formData = new FormData();
    formData.append('file', file);

    axios.post(`${API_URL}upload`, formData, OPTIONS)
      .then(_ => {
        fileInputRef.current.reset()
        console.log('File uploaded successfully');
        setFile(null); // Reset file input after upload
        // Re-fetch files after upload
        axios.get(API_URL, OPTIONS)
          .then(response => setFilesList(response.data.files));
      })
      .catch(error => console.error('Error uploading file:', error));
  };

  return (
    <div>
      <h1>File Upload</h1>
      <form ref={fileInputRef}>
        <input type="file" onChange={handleFileChange} />
        <button onClick={handleFileUpload}>Upload File</button>
      </form>

      <h2>Available Files</h2>
      <ul>
        {filesList.map((file, index) => (
          <li key={file.URL}><a href={file.URL}>{file.Name}</a></li>
        ))}
      </ul>
    </div>
  );
};

export default App;
