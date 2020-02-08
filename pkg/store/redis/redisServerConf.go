package redis

import "strconv"

func (rds *RdsCliBox) GetServerConfig(para ...string) (map[string]string, error) {
	var f map[string]string
	if len(para) > 0 {
		for _, v := range para {
			d, err := rds.conn.ConfigGet(v).Result()
			if err != nil {
				f[v] = (err).Error()
			} else {
				for _, dv := range d {
					f[v] = f[v] + dv.(string) + ","
				}
			}
		}
	} else {
		d, err := rds.conn.ConfigGet("*").Result()
		if err != nil {
			f["*"] = (err).Error()
		} else {
			for k, dv := range d {
				f[strconv.Itoa(k)] = dv.(string)
			}
		}
	}
	return f, nil
}

func (rds *RdsCliBox) SetServerConfig(para map[string]string) error {
	for k, v := range para {
		_, err := rds.conn.ConfigSet(k, v).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
