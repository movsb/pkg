package notify

import "testing"

func TestChanify(t *testing.T) {
	c := NewOfficialChanify(`COS96O4GEiJBRE5RRDY1VUFQRklXM1dNWUZTR0JBV0NWVElETlZHTDVJIgcIAhoDbW9u.2Ho5RAqEnCwgciMoAYrduz4cQBZVKca6sJFyYB0Rlg4`)
	c.Send(`[FIRING:2] 证书过期时间 all (10.1.119.164:80 pods) **Firing** test`, `Value: A=6.007327785177419, C=1
	Labels:
	 - alertname = 证书过期时间
	 - addr = blog.twofei.com:443
	 - grafana_folder = all
	 - instance = 10.1.119.164:80
	 - job = pods
	Annotations:
	 - description = 剩余时间：$A 天。
	 - summary = 服务器证书快要过期了！
	Source: /grafana/alerting/grafana/cb535cbe-1c2f-4b5a-bb07-ffe8da1d5a36/view
	Silence: /grafana/alerting/silence/new?alertmanager=grafana&matcher=addr%3Dblog.twofei.com%3A443&matcher=alertname%3D%E8%AF%81%E4%B9%A6%E8
	%BF%87%E6%9C%9F%E6%97%B6%E9%97%B4&matcher=grafana_folder%3Dall&matcher=instance%3D10.1.119.164%3A80&matcher=job%3Dpods`, true)
}
