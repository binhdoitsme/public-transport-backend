// import necessary modules
import { check } from "k6";
import http from "k6/http";

// define configuration
export const options = {
  // define thresholds
  thresholds: {
    http_req_failed: [{ threshold: "rate<0.01", abortOnFail: true }], // availability threshold for error rate
    http_req_duration: ["p(99)<1000"], // Latency threshold for percentile
  },
  // define scenarios
  scenarios: {
    breaking: {
      executor: "ramping-vus",
      stages: [
        { duration: "10s", target: 20 },
        { duration: "50s", target: 20 },
        { duration: "50s", target: 40 },
        { duration: "50s", target: 60 },
        { duration: "50s", target: 80 },
        { duration: "50s", target: 100 },
        { duration: "50s", target: 120 },
        { duration: "50s", target: 140 },
        //....
      ],
    },
  },
};

export default function () {
  // define URL and request body
  const url = "http://127.0.0.1:8000/api/v1/passengers";
  const payload = JSON.stringify({
    phoneNumber: "+84383462194",
    vneId: "VNE123456789",
    name: "John Doe",
    dob: "1990-01-01T00:00:00Z",
    gender: "Male",
    personalImage: "https://example.com/images/johndoe.png",
    accountType: "Individual",
  });
  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  // send a post request and save response as a variable
  const res = http.post(url, payload, params);

  // check that response is 200
  check(res, {
    "response code was 200": (res) => res.status == 200,
  });
}
