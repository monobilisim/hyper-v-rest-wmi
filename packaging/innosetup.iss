[Setup]
AppName=HyperV REST WMI
AppVersion=2.1
AppPublisher=Mono Bilisim
AppPublisherURL=https://mono.net.tr
AppSupportURL=https://github.com/monobilisim/hyper-v-rest-wmi
AppUpdatesURL=https://github.com/monobilisim/hyper-v-rest-wmi
DefaultDirName={code:GetProgramFiles}\hyper-v-rest-wmi
UsePreviousAppDir=false
UninstallDisplayIcon={app}\hyper-v-rest-wmi.exe
OutputBaseFilename=hyper-v-rest-wmi-setup
Compression=lzma
SolidCompression=yes

[Code]
function GetProgramFiles(Param: string): string;
begin
  if IsWin64 then Result := ExpandConstant('{commonpf64}')
    else Result := ExpandConstant('{commonpf32}')
end;

procedure TaskKill(FileName: String);
var
  ResultCode: Integer;
begin
    Exec('taskkill.exe', '/f /im ' + '"' + FileName + '"', '', SW_HIDE,
     ewWaitUntilTerminated, ResultCode);
end;


[Files]
Source: "hyper-v-rest-wmi.exe"; DestDir: "{app}"; BeforeInstall: TaskKill('hyper-v-rest-wmi.exe')
Source: "cleanup.bat"; DestDir: "{app}"


[Icons]
;Name: "{group}\HyperV REST WMI"; Filename: "{app}\hyper-v-rest-wmi.exe"
;Name: "{group}\Uninstall"; Filename: "{uninstallexe}"


[Messages]
SetupAppTitle=HyperV REST WMI


[Run]
Filename: "{app}\cleanup.bat"
Filename: "{app}\hyper-v-rest-wmi.exe"; Description: "Install Service"; Parameters: --service=install
Filename: "{app}\hyper-v-rest-wmi.exe"; Description: "Start Service"  ; Parameters: --service=start


[UninstallRun]
Filename: "{app}\cleanup.bat"
Filename: "{app}\hyper-v-rest-wmi.exe"; Parameters: --service=stop
Filename: "{app}\hyper-v-rest-wmi.exe"; Parameters: --service=uninstall
Filename: "{cmd}"; Parameters: "/C ""taskkill /im hyper-v-rest-wmi.exe /f /t"
