package httpx_sample

import "httpx"

nums: [1,2,3]

headers: {
	for i, v in nums {
	  "\(i)": ["\(i)"]
  }
}

resp: httpx.JsonGet("https://6544a6335a0b4b04436ca69e.mockapi.io/api/v1/sample/do-request", headers, headers)