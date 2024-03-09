// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyB7ipnjRbNv2xe-T_2WQDJxh0ffesb4MHs",
  authDomain: "trinit-httpsms.firebaseapp.com",
  projectId: "trinit-httpsms",
  storageBucket: "trinit-httpsms.appspot.com",
  messagingSenderId: "15327853725",
  appId: "1:15327853725:web:98cf0940f136e3105ae874",
  measurementId: "G-D261LNGRNT"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);