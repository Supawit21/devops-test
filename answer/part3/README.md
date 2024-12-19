# Design and Troubleshoot Monitoring
**เขียนแผนตรวจสอบเพื่อระบุสาเหตุ โดยใช้เครื่องมือเช่น Prometheus, Grafana หรือ ELK Stack**
----------
| Description | Focus Area | Tools | 
| --- | --- | --- |
| การใช้ CPU หรือ Memory ใน Pod สูงเกินไป | CPU Usage and Memory Usage (Node/Pod) | Prometheus,Grafana |
| มี Error ใน API Logs | Activity Logs Message | Elasticsearch,Kibana |
| Latency ของเครือข่ายสูง | Node Network Traffic | Prometheus,Grafana |
----------
**เสนอวิธีแก้ปัญหา**
----------
| Description | Suggestions |
| --- | --- |
| การใช้ CPU หรือ Memory ใน Pod สูงเกินไป | Scale Horizontal Pod Autoscaler (HPA) and Vertical Pod Autoscaling (VPA) ให้กับ pod เพื่อลดการใช้งานของ CPU, Memory |
| มี Error ใน API Logs | ทำ Alert เพื่อแจ้งเตือน Error API Logs |
| Latency ของเครือข่ายสูง | ใช้งาน Load Balancing เพื่อให้กระจาย Traffic ระหว่าง Pod |
----------
**แนะนำแนวทางป้องกันปัญหาลักษณะนี้ในอนาคต**
----------
| Description | Suggestions |
| --- | --- |
| การใช้ CPU หรือ Memory ใน Pod สูงเกินไป | ปรับปรุง application และการ Query ข้อมูลจาก Database เพื่อลดการใช้งานทั้ง CPU, Memory |
| มี Error ใน API Logs | ทำ Intregrate Test เพื่อตรวจสอบ API ก่อนที่จะทำการ Update To Production |
| Latency ของเครือข่ายสูง | ใช้ Service Mesh เช่น Istio เพื่อมาช่วย traffic routing and retries. |