package repository

const (
	addNewPath = `INSERT INTO default.visit_history (
						short_path, domain, created_at,
				        latitude, longitude, country,
				        city, proxy
					) VALUES (
						$1, $2, $3,
						$4, $5, $6,
				    	$7, $8
					);`

	addPathVisit = `INSERT INTO default.visit_history (
						short_path, visited_at,
				        latitude, longitude, country,
				        city, proxy
					) VALUES (
						$1, $2, 
						$3, $4, $5, 
					    $6, $7
					);`

	listPathVisits = `	SELECT	visited_at, latitude, longitude,
       							country, city, proxy
						FROM default.visit_history
						WHERE short_path = $1 AND visited_at > 0;`
)
