# Cat-API

A API that Queries a specific cat breed

## How to install it

- git clone $project
- cd $project
- docker-compose -f "docker-compose.yml" up -d --build
- the API will be available on port :8081 and the database on port :3306

## Login

- The current Endpoint for login will be available on: POST <http://localhost:8081/login>
- The credentials needs to be passed as json on the request body (these credentials are real and will work):

```json
{
 "password":"@#$RF@!718",
 "username":"admin"
}
```

- if successful the response will be:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNTc4MDQxMjIzfQ.M1tjT6G0ICkvhMHKkuI_OO0N5PxwJqNz2_rc-4qIfxM"
}
```

- This token that needs to be passed in the breeds

## Breeds

- The current Endpoint for getBreeds will be available on: GET <http://localhost:8081/breeds?name=id_of_breed>
- The Token needs to be passed on the header with key `Bearer-Token: token`
- If successful the response will be something like:

```json
{
  "breed": {
    "adaptability": 5,
    "affection_level": 5,
    "alt_names": "",
    "cfa_url": "http://cfa.org/Breeds/BreedsAB/Abyssinian.aspx",
    "child_friendly": 3,
    "country_code": "EG",
    "country_codes": "EG",
    "description": "The Abyssinian is easy to care for, and a joy to have in your home. They’re affectionate cats and love both people and other animals.",
    "dog_friendly": 4,
    "energy_level": 5,
    "experimental": 0,
    "grooming": 1,
    "hairless": 0,
    "health_issues": 2,
    "hypoallergenic": 0,
    "id": "abys",
    "indoor": 0,
    "intelligence": 5,
    "lap": 1,
    "life_span": "14 - 15",
    "name": "Abyssinian",
    "natural": 1,
    "origin": "Egypt",
    "rare": 0,
    "rex": 0,
    "shedding_level": 2,
    "short_legs": 0,
    "social_needs": 5,
    "stranger_friendly": 5,
    "suppressed_tail": 0,
    "temperament": "Active, Energetic, Independent, Intelligent, Gentle",
    "vcahospitals_url": "https://vcahospitals.com/know-your-pet/cat-breeds/abyssinian",
    "vetstreet_url": "http://www.vetstreet.com/cats/abyssinian",
    "vocalisation": 1,
    "weight": {
      "imperial": "7  -  10",
      "metric": "3 - 5"
    },
    "wikipedia_url": "https://en.wikipedia.org/wiki/Abyssinian_(cat)"
  }
}
```

## Contact

- If you have any questions just send me a email at felipe.pbgomes@gmail.com
