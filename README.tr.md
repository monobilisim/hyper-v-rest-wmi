[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![GPL License][license-shield]][license-url]

[![Readme in English](https://img.shields.io/badge/Readme-English-blue)](README.md)
[![Readme in Turkish](https://img.shields.io/badge/Readme-Turkish-red)](README.tr.md)

<div align="center">
<a href="https://mono.net.tr/">
  <img src="https://monobilisim.com.tr/images/mono-bilisim.svg" width="340"/>
</a>

<h2 align="center">hyper-v-rest-wmi</h2>
<b>hyper-v-rest-wmi</b> Hyper-V sanal makine bilgilerine RESTful API aracılığıyla erişim sağlamak için bir araçtır.
</div>

---

## İçindekiler

- [İçindekiler](#i̇çindekiler)
- [Kurulum](#kurulum)
- [Kullanım](#kullanım)
- [Lisans](#lisans)

---

## Kurulum

1. [Releases](https://github.com/monobilisim/hyper-v-rest-wmi/releases) sayfasına gidin.
2. *Portable (taşınabilir, kurulumsuz)* veya *Setup (kurulumlu)* paketlerden birini indirin.
3. Çalıştırın!

## Kullanım

- Tüm sanal makineler: `/vms`
- CPU bilgileri: `/vms/<VM-ID>/summary`
- RAM bilgileri: `/vms/<VM-ID>/memory`
- Disk bilgileri: `/vms/<VM-ID>/vhd`
- Ağ bilgileri: `/vms/<VM-ID>/ip`
- Sürüm: `/version`

---

## Lisans

hyper-v-rest-wmi, GPL-3.0 lisanslıdır. Detaylar için [LICENSE](LICENSE) dosyasına bakınız.

[contributors-shield]: https://img.shields.io/github/contributors/monobilisim/hyper-v-rest-wmi.svg?style=for-the-badge
[contributors-url]: https://github.com/monobilisim/hyper-v-rest-wmi/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/monobilisim/hyper-v-rest-wmi.svg?style=for-the-badge
[forks-url]: https://github.com/monobilisim/hyper-v-rest-wmi/network/members
[stars-shield]: https://img.shields.io/github/stars/monobilisim/hyper-v-rest-wmi.svg?style=for-the-badge
[stars-url]: https://github.com/monobilisim/hyper-v-rest-wmi/stargazers
[issues-shield]: https://img.shields.io/github/issues/monobilisim/hyper-v-rest-wmi.svg?style=for-the-badge
[issues-url]: https://github.com/monobilisim/hyper-v-rest-wmi/issues
[license-shield]: https://img.shields.io/github/license/monobilisim/hyper-v-rest-wmi.svg?style=for-the-badge
[license-url]: https://github.com/monobilisim/hyper-v-rest-wmi/blob/master/LICENSE
