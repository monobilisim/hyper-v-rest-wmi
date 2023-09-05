## Kurulum

- `%PROGRAMFILES%\hyper-v-rest` klasörü oluşturulur.
- `hyper-v-rest.exe` dosyası, `%PROGRAMFILES%\hyper-v-rest` klasörüne kopyalanır.
- Windows PowerShell, yönetici olarak açılır ve aşağıdaki komutlar çalıştırılır:

       cd "$env:PROGRAMFILES\hyper-v-rest"
       .\hyper-v-rest.exe --service=install
       .\hyper-v-rest.exe --service=start

## Kullanım

- Tüm VM'ler: `/vms`
- Processor sayısı ve GuestOS: `/vms/<Name>/summary`
- Memory: `/vms/<Name>/memory`
- Disk bilgisi `/vms/<Name>/vhd`
- IP numaraları `/vms/<Name>/ip`
- Versiyon: `/version`
