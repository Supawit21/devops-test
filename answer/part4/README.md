# Question and Answer
**Service Mesh และ API Gateway**
----------
1. อธิบายความแตกต่างระหว่าง Istio และ Kong
    - **Istio** เป็น Service Mesh ที่รองรับการสื่อสาร service ต่างๆ ใน microservices โดยแบ่งคุณสมบัติ 4 อย่าง Connect,Secure,Control,Observe
    - **Kong** เป็น API Gateway และ API Management Tools เป็นตัวกลางในการสื่อสารระหว่าง Clients กับ Microservice ต่างๆ ช่วยบริหารจัดการ Functionality ทั้งหมดเลยไม่ว่าจะเป็น Caching, Monitoring, Authentication หรือแม้กระทั้ง Rate-limiting

2. ยกตัวอย่างการใช้งานทั้งสองร่วมกัน
    - ใช้ Kong ในการสื่อสารระหว่าง Client กับ Microservice ที่มีและใช้ Istio ช่วยในการรองรับการสื่อสารระหว่าง service หรือตรวจสอบการ connect ระหว่าง service หรือเรื่องการจัดการความปลอดภัย policy ต่างๆและการ Observe ภายใต้ Cluster นั้น
----------
**Security และ Compliance**
----------
3. คุณจะป้องกันความปลอดภัยใน CI/CD Pipeline ได้อย่างไร?
    - เก็บข้อมูล sensitive data ที่ใช้ในการทำ CI/CD Pipeline ไว้ในรูปแบบ Credential หรือ Secrets เพื่อสามารถจัดการได้ (secrets management)
4. คุณจะจัดการ Kubernetes Secrets อย่างปลอดภัยได้อย่างไร?
    - อาจจะมีการใช้ third party เพื่อมาเก็บ secrets ใน kubernetes เช่น Hashicorp Vault ในกรณีที่สร้าง cluster on Production
----------
**Security และ Compliance**
----------
5. คุณจะออกแบบ Kubernetes Cluster บน AWS ให้รองรับ High Availability ได้อย่างไร?
    - ควรมีการแบ่ง site สำรองเพื่อรองรับในกรณีที่ site หลักเกิดปัญหาและสามารถ switch ไปยัง site ที่สำรองไว้ได้ DC site --> DR site
    - ทำ Load Balancing เพื่อรองรับและกระจาย traffic
    - ทำ Auto Scaling ในระดับ node จนไปถึงการทำ Auto Scaling ในระดับ pod
6. อธิบายข้อดีของการใช้ Horizontal Pod Autoscaling ใน Kubernetes
    - สามารถทำให้เพิ่มหรือลดจำนวน pod ได้แบบอัติโนมัติ
    - สามารถรองรับ metrics ที่สามารถกำหนดขึ้นมาเองได้เช่น Request